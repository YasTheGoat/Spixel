package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Spy struct {
	Id     uuid.UUID
	Active bool
	Read   bool
	Time   time.Time
	IP     string
	Name   string
}

var targets map[string]Spy

func main() {

	router := chi.NewRouter()

	fs := http.FileServer(http.Dir("."))

	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		val, ok := targets[id]
		if ok {
			if val.Active {
				targets[id] = Spy{
					Id:     val.Id,
					Active: false,
					Read:   true,
					Time:   time.Now(),
					IP:     ReadUserIP(r),
					Name:   val.Name,
				}

				SaveToFile()
			}
			fs.ServeHTTP(w, r)
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}

	})

	log.Println("Server started on port : 8080")
	ReadFromFile()
	go CLI()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalln("Failed to start server")
	}
}

func CLI() {

	for {
		fmt.Println("What do you want to do ?")
		fmt.Println("1. List active targets")
		fmt.Println("2. List read targets")
		fmt.Println("3. List all targets")
		fmt.Println("4. Add target")
		fmt.Println("5. Activate target")
		fmt.Print("Choice : ")

		var input string
		fmt.Scanln(&input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter a valid choice please.")
			continue
		}

		switch choice {
		case 1:
			count := 0
			for _, value := range targets {
				if value.Active {
					count++
					fmt.Println(count, ". ID:", value.Id, " Name:", value.Name, " Active:", value.Active)
				}
			}
			if count == 0 {
				fmt.Println("No target found")
			}
		case 2:
			count := 0
			for _, value := range targets {
				if value.Read {
					count++
					fmt.Println(count, ". ID:", value.Id, " Name:", value.Name, " Active:", value.Active, " Read:", value.Read, " Time:", value.Time)
				}
			}
			if count == 0 {
				fmt.Println("No target found")
			}
		case 3:
			count := 0
			for _, value := range targets {
				count++
				fmt.Println(count, ". ID:", value.Id, " Name:", value.Name, " Active:", value.Active, " Read:", value.Read, " Time:", value.Time)
			}
			if count == 0 {
				fmt.Println("No target found")
			}
		case 4:
			fmt.Print("Choose name : ")
			var name string
			fmt.Scanln(&name)

			id := uuid.New()
			targets[id.String()] = Spy{
				Id:     id,
				Name:   name,
				Active: false,
				Read:   false,
			}
		case 5:
			fmt.Print("Enter the id of the target to activate : ")
			var id string
			fmt.Scanln(&id)

			val, ok := targets[id]
			if !ok {
				fmt.Println("Target not found")
				continue
			}

			targets[id] = Spy{
				Id:     val.Id,
				Name:   val.Name,
				Active: true,
				Read:   false,
			}
		default:
			fmt.Println("Invalid choice ! Try again.")
			continue
		}

		fmt.Print("\n\n")
		SaveToFile()
	}

}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func SaveToFile() {
	file, err := os.Create("save.txt")
	if err != nil {
		fmt.Println("Failed to open/create save file")
		return
	}

	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)

	if err := e.Encode(targets); err != nil {
		fmt.Println("Failed to save targets to file")
		return
	}

	file.Write(b.Bytes())
}

func ReadFromFile() {
	data, err := os.ReadFile("save.txt")
	if err != nil {
		fmt.Println("Failed to read from save file")
		return
	}

	b := new(bytes.Buffer)
	b.Write(data)

	d := gob.NewDecoder(b)

	if err := d.Decode(&targets); err != nil {
		fmt.Println("Failed to retrive saved data")
		return
	}

}
