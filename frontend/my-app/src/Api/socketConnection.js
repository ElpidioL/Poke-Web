import { Intent } from "../scripts/intentVerify.js";
// api/index.js
var socket = new WebSocket("ws://localhost:8080/ws");

let Connect = () => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

   socket.onmessage = msg => {
    //console.log(JSON.stringify(msg.data).replace(/\\/g, ''))
    Intent(JSON.parse(msg.data.replace(/\\/g, '')))
  }; 

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };  
};

let Send =(info) =>{
  socket.send(JSON.stringify(info));
}
export { Connect, Send };