NAME = phonebook
SRCS = main.go
SRCS += contact.go phonebook.go 
CC 	 = go build

all:$(NAME)

$(NAME):
		$(CC) -o $(NAME) $(SRCS) 

clean: 
		go clean

fclean: clean
		rm -rf $(NAME)

re: fclean all

.PHONY: all clean fclean re