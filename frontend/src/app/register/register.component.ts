import { Component, OnInit } from '@angular/core';
import { UserService } from '../services/user.service';
import { register } from '../data-type';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit{
  title = 'NDrive Challenge';

  constructor(private user:UserService){}

  ngOnInit(): void{
    this.user.reloadUser()
  }
  register(data:register){
    this.user.userRegister(data)
  } 
}
