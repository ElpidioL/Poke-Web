import { useState, useEffect } from "react";
import classes from "./IndexPage.module.css";
import { Colour,PokemonNew } from "../Defaults/classes"
import { Connect,Send } from "../Api/socketConnection";
import { GetEmail, GetColour, GetCredits, GetInfo, GetUpdate } from "../scripts/getCookies";

function GetPokemon(){
/*   let a = `{"pokemons": [], "toHatch": {"10": 20}}`
  let b = JSON.parse(a)
  let c = Object.keys(b.toHatch);
  console.log(typeof( c[0])) */
  let pokemon = new PokemonNew();
  let credits = GetCredits();
  if(credits >= 0){
    document.cookie = `Credits=${credits - 30}`;
    fetch('https://pokeapi.co/api/v2/pokemon')
    .then(response => response.json())
    .then(data => {
      fetch(`https://pokeapi.co/api/v2/pokemon?limit=1&offset=${Math.floor(Math.random() * (data.count - 1)) + 1}`)
      .then(response => response.json())
      .then(data2 => {
        fetch(`https://pokeapi.co/api/v2/pokemon/${data2.results[0].name}`)
        .then(response => response.json())
        .then(data3 => {
          pokemon.pokemon = data2.results[0].name
          pokemon.pokeId = data3.id
          pokemon.intent = "pokemon"
          Send(pokemon);
        });
      });
    });
  }else{
    console.log("no credits.")
  }
}

function deleteCookies(){
  document.cookie = "Email=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
  document.cookie = "Colour=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
  document.cookie = `Credits=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
  document.cookie = `Info=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
  document.cookie = `Last=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
  function Nav() {
    window.location.reload();
  }
  setTimeout(Nav, 200); 
}


function IndexPage(){
  const [isLoading, setIsLoading] = useState(true);
  const [credits, setCredits] = useState(0);
  const [info, setInfo] = useState("");
  const [lastUpdate, setLastUpdate] = useState("");
  
 

  useEffect(() => {
    function LoadConfig() {
      Connect();
      let colour = new Colour(GetColour(), GetEmail(), "colour")
      Send(colour);
    }
    setTimeout(LoadConfig, 50);
    const timer = setTimeout(() => {
      setIsLoading(true)
      setCredits(GetCredits());
      setInfo(GetInfo());
      setLastUpdate(GetUpdate());
      if(GetEmail() !== ""){
        setIsLoading(false)
      }
    }, 800);
    return () => {clearTimeout(timer);}
  }, []);


    if(isLoading){
      return (
        <div className={classes.IndexPage}>
          <h1>Loading...</h1>
          <button onClick={deleteCookies}>clear</button>
        </div>
        )
    }else{
        return (
          <div className={classes.IndexPage}>
              <h1>Hello {GetEmail()}</h1>
              <p>Credits: {credits}, <br/> Info: {info},  <br/> Last update: {lastUpdate}</p>
              <button onClick={() => {GetPokemon()}}>poke</button>
          </div>
        )
      }
  }


export default IndexPage;

