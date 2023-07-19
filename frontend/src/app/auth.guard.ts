import { CanActivateFn } from '@angular/router';
import { UserService } from './register/user.service';
import { inject } from '@angular/core';

export const authGuard: CanActivateFn = (route, state) => {
  console.warn(localStorage.getItem('user'))
  if(localStorage.getItem('user') != null){
    console.warn("trueee")
    return false;
  }
  return true;
};
