import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegisterComponent } from './register/register.component';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';
import { authGuard, isLogged } from './auth.guard';
import { AddProductComponent } from './add-product/add-product.component';
import { UpdateProductComponent } from './update-product/update-product.component';
import { ProductPageComponent } from './product-page/product-page.component';
import { SearchPageComponent } from './search-page/search-page.component';

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
  {path: 'home', component: HomeComponent},
  {path:'add-product', component:AddProductComponent, canActivate:[isLogged]},
  {path:'update-product/:id', component:UpdateProductComponent, canActivate:[isLogged]},
  {path:'product/:id', component:ProductPageComponent},
  {path:'search/:query', component:SearchPageComponent}
  // Add other top-level routes if required
];

@NgModule({
  imports: [RouterModule.forRoot(routes,{
    onSameUrlNavigation:'reload'
  })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
