import { Intent } from "../scripts/intentVerify.js";
import { Handshake } from "../Defaults/classes"
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

let Send =(info, handshake) =>{
  handshake ? info.handshake = handshake : handshake = ""
  socket.send(JSON.stringify(info));
}
export { Connect, Send };