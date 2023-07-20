import { Component, OnInit } from '@angular/core';
import { category, product } from '../data-type';
import { ProductService } from '../services/product.service';

@Component({
  selector: 'app-add-product',
  templateUrl: './add-product.component.html',
  styleUrls: ['./add-product.component.css']
})
export class AddProductComponent implements OnInit{
  categoryList:category[] = [];
constructor(private productService:ProductService){}

ngOnInit():void{
  this.productService.categoryList().subscribe((res:category[])=>{
    if(res){
      this.categoryList= res
    }
  })
}

submitProduct(data:product){
  console.warn(data)
  this.productService.addProduct(data)
}
}
