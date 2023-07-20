import { Component } from '@angular/core';
import { product } from '../data-type';
import { ProductService } from '../services/product.service';
import { Router } from '@angular/router';
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  
  isLoggedIn=false

  productList:product[] = [];
  constructor(private productService:ProductService, private router:Router){}

ngOnInit():void{
  this.productService.productList().subscribe((res:product[])=>{
    if(res){
      this.productList= res
    }
  })
}

deleteProduct(id:string){
  this.productService.deleteProduct(id).subscribe((res)=>{
    if(res){
      console.warn("product deleted with ID", id)
      this.router.routeReuseStrategy.shouldReuseRoute = () => false;
      this.router.onSameUrlNavigation = 'reload';
      this.router.navigate(['./']),{
        relativeTo: this.router
      }
    }
  })
}
}
