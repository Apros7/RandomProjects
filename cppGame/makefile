CFLAGS = -std=c++11 -Wall -pedantic

bin/main: bin/main.o bin/game.o bin/ObjectField.o
	g++ -g -o bin/main bin/main.o bin/game.o bin/ObjectField.o $(CFLAGS) -lncurses

bin/main.o: src/main.cpp
	mkdir -p bin
	g++ -g -c -o bin/main.o src/main.cpp

bin/game.o: src/game.cpp src/game.h
	mkdir -p bin	
	g++ -g -c -o bin/game.o src/game.cpp $(CFLAGS)

bin/ObjectField.o: src/ObjectField.cpp src/ObjectField.h
	mkdir -p bin	
	g++ -g -c -o bin/ObjectField.o src/ObjectField.cpp $(CFLAGS)


run: bin/main
	./bin/main