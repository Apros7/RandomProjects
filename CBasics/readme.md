gcc -o output_file src_file
./outputfile

to fix leaks on mac: leaks --atExit -- ./YOUR_PROGRAM_NAME
make sure to -g in gcc before.

to get error code use perror