slint::include_modules!();
use chrono::Datelike;
// use native_dialog::{FileDialog, MessageDialog, MessageType};
use native_dialog::{FileDialog};
use std::fs;
use std::path::PathBuf;

fn main() -> Result<(), slint::PlatformError> {
    let ui = AppWindow::new()?;

    ui.on_request_update_date({
        let ui_handle = ui.as_weak(); 
        move || {
            let ui = ui_handle.unwrap();
            let current_month = chrono::Local::now().month();
            ui.set_month(match current_month {
                1 => String::from("January"),
                2 => String::from("February"),
                3 => String::from("March"),
                4 => String::from("April"),
                5 => String::from("May"),
                6 => String::from("June"),
                7 => String::from("July"),
                8 => String::from("August"),
                9 => String::from("September"),
                10 => String::from("October"),
                11 => String::from("November"),
                12 => String::from("December"),
                _ => String::from("Unknown"),
            }.into());
            let current_day = chrono::Local::now().day();
            ui.set_day(current_day.try_into().unwrap());
        }
    });

    ui.on_request_upload_file({
        let ui_handle = ui.as_weak(); 
        move || {
            let mut message = String::new();
            let ui = ui_handle.unwrap();
            let path = FileDialog::new()
                .set_location("~/Desktop")
                .add_filter("CSV csv", &["csv"])
                .show_open_single_file()
                .unwrap();
            if let Some(path) = path {
                let file_name = path.file_name().unwrap().to_string_lossy().into_owned();
                let new_path = PathBuf::from("./data").join(&file_name);
    
                // Attempt to copy the file to the data directory
                if let Err(e) = fs::copy(&path, &new_path) {
                    message = e.to_string()
                } else {
                    message = String::from("Success")
                }
            } else {
                message = String::from("No file selected");
            }
            ui.set_month(String::from(message).into())
        }
    });
        

        // use native_dialog::{FileDialog, MessageDialog, MessageType};
        // let path = FileDialog::new()
        // .set_location("~/Desktop")
        // .add_filter("PNG Image", &["png"])
        // .add_filter("JPEG Image", &["jpg", "jpeg"])
        // .show_open_single_file()
        // .unwrap();

        // let path = match path {
        //     Some(path) => path,
        //     None => return,
        // };

        // let yes = MessageDialog::new()
        // .set_type(MessageType::Info)
        // .set_title("Do you want to open the file?")
        // .set_text(&format!("{:#?}", path))
        // .show_confirm()
        // .unwrap();

        // if yes {
        //     // do_something(path);
        // };


    ui.run()    
}
