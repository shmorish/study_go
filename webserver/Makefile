all :
	@echo "Usage: make [ install / clean / re / ex~ ]"

install : 
	@echo "Installing webserver..."
	@go mod init webserver 2> /dev/null || echo "webserver already initialized."
	@go get github.com/labstack/echo/v4 2> /dev/null
	@echo "Done."

clean :
	@echo "Cleaning webserver..."
	@rm go.mod go.sum > /dev/null 2> /dev/null || echo "No go.mod or go.sum file found."
	@rm copied_avatar.jpg > /dev/null 2> /dev/null || echo "No copied_avatar.jpg file found."
	@echo "Done."

re : clean install

ex00 :
	@echo "Running example 00..."
	@echo "Open http://localhost:8080/ in your browser."
	@go run ex00/main.go || make re > /dev/null

ex01 :
	@echo "Running example 01..."
	@echo "Open http://localhost:8080/users/joe in your browser."
	@echo "Open http://localhost:8080/show?team=Japan&member=Shoto in your browser."
	@go run ex01/main.go || make re > /dev/null

ex02 :
	@echo "Running example 02..."
	@echo "Run the following commands in your terminal:"
	@echo "curl -d \"name=First Family\" -d \"email=Shoto@university.com\" http://localhost:8080/save"
	@go run ex02/main.go || make re > /dev/null

ex03 :
	@echo "Running example 03..."
	@echo "Run the following commands in your terminal:"
	@echo "curl -F "name=John" -F \"avatar=@$(PWD)/ex03/gopher-ninja.png\" http://localhost:8080/save"
	@go run ex03/main.go || make re > /dev/null

.PHONY: all install clean re ex00 ex01 ex02 ex03
