package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>May Challenge</title>
		<style>
      *,::before,::after{
        margin: 0;
        padding: 0;
        box-sizing: border-box;
      }
      
      header {
				margin-left: 35%;
      }

      h2 {
        font-size: 1.6rem;
        font-weight: 300;
         color: #15133C;
       }

      h1{
          font-size: 2.4rem;
          font-weight: 600;
          color: #15133C;
        }

			div {
        border: 3px outset rgb(63, 63, 63);
				justify-content: center;
				display: flex;
       
			}
		</style>
	</head>
	<body style="background: #E8F9FD;">
    <header>
      <h1>Take control of yourself</h1>
    </header>
		<div>
      <h2
				Out of the night that covers me <br />

				Black as the pit from pole to pole <br />

				I thank whatever gods may be <br />

				For my unconquerable soul. <br />

				In the fell clutch of circumstance <br />

				I have not winced nor cried aloud. <br />

				Under the bludgeonings of chance <br />

				My head is bloody, but unbowed. <br />

				Beyond this place of wrath and tears <br />

				Looms but the Horror of the shade, <br />

				And yet the menace of the years <br />

				Finds and shall find me unafraid. <br />

				It matters not how strait the gate, <br />

				How charged with punishments the scroll, <br />

				I am the master of my fate, <br />

				I am the captain of my soul.
			</h2>
		</div>
	</body>
</html>
`)
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on 5000....")
	http.ListenAndServe(":5000", nil)
}
