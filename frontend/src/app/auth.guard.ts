import { CanActivateFn } from '@angular/router';
import { UserService } from './services/user.service';
import { inject } from '@angular/core';

export const authGuard: CanActivateFn = (route, state) => {
  console.warn(localStorage.getItem('user'))
  if(localStorage.getItem('user') != null){
    console.warn("trueee")
    return false;
  }
  return true;
};

export const isLogged: CanActivateFn = (route, state) => {
  console.warn(localStorage.getItem('user'))
  if(localStorage.getItem('user') != null){
    return true;
  }
  return false;
};
