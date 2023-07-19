import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../register/user.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit{
  menuType:String = 'default'
  constructor(private userService:UserService) {
    
  }
ngOnInit(): void {
  if(localStorage.getItem('user')){
    this.menuType="user"
  }else{
    this.menuType="default"
  }
}
logout(){
  this.userService.userLogout()
} 
}
