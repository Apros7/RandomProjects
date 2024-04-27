#include <unistd.h>
#include <ncurses.h>

#include <cstdint>
#include <string>
#include <vector>

#include "game.h"
#include "ObjectField.h"

WINDOW* main_wnd;
WINDOW* game_wnd;

rect game_area;
rect screen_area;

vec2ui cur_size;

ObjectField asteroids;
ObjectField stars;

struct {
    vec2i pos;
    rect bounds;
    char disp_char;
    int energy;
} player;

int init() {
    srand(time(0));

    main_wnd = initscr();
    cbreak();
    noecho();
    clear();
    refresh();

    curs_set(0);

    start_color();

    screen_area = { { 0, 0 }, { 80, 24 } };

    int infopanel_height = 4;
    game_wnd = newwin(  screen_area.height() - infopanel_height - 2, 
                        screen_area.width() - 2, 
                        screen_area.top() + 1, 
                        screen_area.left() + 1  );
    main_wnd = newwin(screen_area.height(), screen_area.width(), 0, 0);

    game_area = { { 0, 0 }, { static_cast<int_fast8_t>(screen_area.width() - 2), static_cast<int_fast8_t>(screen_area.height() - infopanel_height - 4) } };

    init_pair(1, COLOR_WHITE, COLOR_BLACK);
    init_pair(2, COLOR_GREEN, COLOR_BLACK);
    init_pair(3, COLOR_YELLOW, COLOR_BLACK);
    init_pair(4, COLOR_RED, COLOR_BLACK);
    init_pair(5, COLOR_BLUE, COLOR_BLACK);
    
    // enable function keys
    keypad(main_wnd, true);
    keypad(game_wnd, true);
   
    // disable input blocking
    nodelay(main_wnd, true);
    nodelay(game_wnd, true);

    // test for color here

    return 0;
}


void run() {

    int tick;

    // initialize player as before

    // constrain object fields to game area
    asteroids.setBounds(game_area);
    stars.setBounds(game_area);

    int in_char = 0;
    bool exit_requested = false;
    bool game_over = false;

    wattron(main_wnd, A_BOLD);
    box(main_wnd, 0, 0);
    wattroff(main_wnd, A_BOLD);

    // draw dividing line between game and stats
    wmove(main_wnd, game_area.bot() + 3, 1);
    whline(main_wnd, '-', screen_area.width() - 2);

    // initial draw
    wrefresh(main_wnd);
    wrefresh(game_wnd);

    tick = 0;
    while(1) {

        // clear game window
        werase(game_wnd);

        // read inputs, lowercase all characters
        in_char = wgetch(main_wnd);
        in_char = tolower(in_char);

        switch(in_char) {
            case 'q':
                exit_requested = true;
                break;
            case KEY_UP:
            case 'w':
            case 'i':
                if(player.pos.y > game_area.top())
                    player.pos.y -= 1;
                break;
            case KEY_DOWN:
            case 's':
            case 'k':
                if(player.pos.y < game_area.bot() + 1)
                    player.pos.y += 1;
                break;
            case KEY_LEFT:
            case 'a':
            case 'j':
                if(player.pos.x > game_area.left() + 1)
                    player.pos.x -= 1;
                break;
            case KEY_RIGHT:
            case 'd':
            case 'l':
                if(player.pos.x < game_area.right() - 2)
                    player.pos.x += 1;
                break;
            default:
                break;
        }        
        if(tick % 7 == 0)
            stars.update();

        for(auto s: stars.getData())
            mvwaddch(game_wnd, s.getPos().y, s.getPos().x, '.');

        if(tick > 100 && tick % 20 == 0)
            asteroids.update();

        /** stars draw **/

        for(auto a : asteroids.getData()) {
            wattron(game_wnd, A_BOLD);
            mvwaddch(game_wnd, a.getPos().y, a.getPos().x, '*');
            wattroff(game_wnd, A_BOLD);
        }
        
        // player ship main body
        wattron(game_wnd, A_BOLD);
        mvwaddch(game_wnd, player.pos.y, player.pos.x, player.disp_char);
        wattroff(game_wnd, A_BOLD);

        wrefresh(main_wnd);
        wrefresh(game_wnd);

        if(exit_requested || game_over) break;

        tick++;

        usleep(10000); // 10 ms
    };
}

void close() {
    endwin();
}