import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../services/user.service';
import { ProductService } from '../services/product.service';
import { product } from '../data-type';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit{
  productList:product[] = [];
  menuType:String = 'default'
  constructor(private userService:UserService, private productService:ProductService) {
    
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

search(data:string){
  console.warn(data)
  this.productService.searchProduct(data).subscribe((res)=>{
    this.productList = res
  })
}
}
