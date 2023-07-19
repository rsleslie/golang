NAME	= 	megaphone
SRCS	=	main.go
OBJS	=	$(SRCS:.go=.o)
CC		=	go build

all:
	$(CC) -o $(NAME)

clean:
	go clean

fclean: clean
	rm -rf $(NAME)

re: fclean all

.PHONY: all clean fclean