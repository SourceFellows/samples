from golang:latest

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN mkdir -p /app/cmd
WORKDIR /app/
COPY . .

#Anwendung kompilieren
RUN go build -gcflags="-N -l" cmd/main.go

#Port auf den sich der Debugger verbinden kann
#Angabe in 'startup.sh'
EXPOSE 4040

#Anwendung starten
RUN chmod +x /app/startup.sh
ENTRYPOINT [ "/app/startup.sh" ]