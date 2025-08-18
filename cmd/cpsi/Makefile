.PHONY: default cstart clean lean



default: cstart cpsi
	@printf "Done.\n"

cstart:
	@printf "Compiling..."

cpsi: main.o clargs.o debugstr.o usock.o
	@printf "Done.\n"
	@printf "Linking..."
	@gcc -o cpsi main.o clargs.o debugstr.o usock.o

main.o: main.c clargs.h debugstr.h usock.h
	@gcc -Wall -Werror -c main.c
	@printf "."

clargs.o: clargs.c clargs.h
	@gcc -Wall -Werror -c clargs.c
	@printf "."

debugstr.o: debugstr.c debugstr.h
	@gcc -Wall -Werror -c debugstr.c
	@printf "."

usock.o: usock.c usock.h
	@gcc -Wall -Werror -c usock.c
	@printf "."

clean: lean
	@printf "Cleaning directory.."
	@rm -f cpsi
	@printf "Done.\n"

lean:
	@printf "Leaning directory.."
	@rm -f *.o
	@printf "."
	@rm -f cpsi.out.*
	@printf "Done.\n"