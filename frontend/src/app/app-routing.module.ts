import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegisterComponent } from './register/register.component';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';
import { authGuard } from './auth.guard';

const routes: Routes = [
  {
    path: 'auth',
    children: [
      { path: 'login', component: LoginComponent, canActivate:[authGuard]},
      { path: 'register', component: RegisterComponent, canActivate:[authGuard]},
      // Add more child routes as needed
    ],
  },
  {path: '', component: HomeComponent},
  {path: 'home', component: HomeComponent}
  // Add other top-level routes if required
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
