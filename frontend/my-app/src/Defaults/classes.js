import { GetEmail, GetColour, GetCredits, GetInfo, GetUpdate } from "../scripts/getCookies";

class LoginPerson {
    constructor(email, password, intent) {
        this.email = email 
        this.password = password
        this.intent = intent
    }
}
class RegisterPerson {
    constructor(email, name, password, intent) {
        this.email = email
        this.name = name 
        this.password = password
        this.intent = intent
    }
}

class Colour {
    constructor(colour, email, intent){
        this.colour = colour
        this.email = email
        this.intent = intent
    }
}


class PokemonNew {
    constructor(pokemon, pokeId, hatchDate, intent) {
      this.pokemon = pokemon
      this.pokeId = pokeId
      this.intent = intent
    }
}


export { LoginPerson, RegisterPerson, Colour ,PokemonNew};