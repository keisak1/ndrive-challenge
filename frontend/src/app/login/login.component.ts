import { Component } from '@angular/core';
import { UserService } from '../register/user.service';
import { login } from '../data-type';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  title = 'NDrive Challenge';
  authError:String= ""
  constructor(private user:UserService){}

  ngOnInit(): void{
    this.user.reloadUser()
  }
  login(data:login){
    this.user.userLogin(data)
    this.user.loginError.subscribe((isError)=>{
      if(isError){
        this.authError="Email or Password are incorrect."
      }
        })
  } 
}
