import { EventEmitter, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { login, register } from '../data-type';
import {Router  } from '@angular/router'
import { BehaviorSubject } from 'rxjs';
import { ProductService } from './product.service';
@Injectable({
  providedIn: 'root',
})
export class UserService {
  isLoggedIn= new BehaviorSubject<boolean>(false)
  loginError = new EventEmitter<boolean>(false)
  constructor(private http: HttpClient, private router:Router, private productService:ProductService) {}

  userRegister(data: register) {
    console.warn('service called');
    return this.http.post(this.productService.url + 'api/auth/register', data, {observe: 'response', withCredentials:true}).subscribe((result) => {
        if (result) {
          localStorage.setItem('user', JSON.stringify(result))
          this.router.navigate(['home'])
        }
      });
  }

  
  userLogin(data:login){

    console.warn('service called');
    return this.http.post(this.productService.url + 'api/auth/login', data, {responseType:'json', withCredentials: true}).subscribe(
      (result) => {
        if (result) {
          
          this.loginError.emit(false)
          localStorage.setItem('user', JSON.stringify(result))
          this.router.navigate(['home'])
        }
      },(error) => { 
        if(error != null){
        this.loginError.emit(true)
        }
      });
  }

  userLogout(){
    return this.http.get(this.productService.url + 'api/auth/logout', {responseType:'json', withCredentials:true}).subscribe(
      (result) => {
        if (result) {
          localStorage.removeItem('user')
          this.router.navigate(['auth/login'])
        }
      },(error) => { 
        console.warn(error)
        if(error != null){
        this.loginError.emit(true)
        }
      });
  }
  reloadUser(){
    if(localStorage.getItem('user')){
      this.isLoggedIn.next(true)
      this.router.navigate(['home'])
    }
  }
}
