import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService, SigninUser } from '@javgat/devtest-api';
import { Mensaje, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';

// SigninComponent es el componente que permite el registro de un nuevo usuario

@Component({
  selector: 'app-signin',
  templateUrl: './signin.component.html',
  styleUrls: ['./signin.component.css']
})
export class SigninComponent implements OnInit {

  signinUserEmpty = {
    username:"",
    email:"",
    pass:""
  }
  // Variable que se modificara en el formulario de registro
  signinUser = this.signinUserEmpty as SigninUser
  mensaje: Mensaje
  
  constructor(private authService : AuthService, private datos : DataService, private router: Router) { 
    this.mensaje = new Mensaje()
  }

  ngOnInit(): void {
    this.datos.mensajeActual.subscribe(
      valor => this.mensaje = valor
    )
  }

  // Envío de petición de registro a BackEnd, y manejo de la respuesta
  signin(su: SigninUser){
    this.authService.registerUser(su).subscribe(
      resp => {        
        this.datos.cambiarMensaje(new Mensaje("Registro con exito", Tipo.SUCCESS, true))
        console.log("Registro con exito")
        this.router.navigate(['login'])
      },
      err =>{
        let msg: string
        if(err.status >= 500)
          msg = "Error al conectar con el servidor"
        else
          msg = err.error.message
        this.datos.cambiarMensaje(new Mensaje("Error al registrar nuevo usuario: "+msg, Tipo.ERROR, true))
        console.log("Error al registrar nuevo usuario: "+msg)
      }
    )
  }

  // Cuando el formulario se envia, se ejecuta la funcio
  signinSubmit(){
    this.signin(this.signinUser)
  }

}
