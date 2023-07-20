import { Component } from '@angular/core';
import { category, product } from '../data-type';
import { ProductService } from '../services/product.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-product-page',
  templateUrl: './product-page.component.html',
  styleUrls: ['./product-page.component.css']
})
export class ProductPageComponent {
  categoryList:category[] = [];
  product = {} as product
constructor(private productService:ProductService, private route:ActivatedRoute){}

ngOnInit():void{
  let productId = this.route.snapshot.paramMap.get('id')
  productId && this.productService.getProduct(productId).subscribe((data:product)=>{
    this.product = data
    console.warn(this.product.categoryId)
  })
  this.productService.categoryList().subscribe((res:category[])=>{
    if(res){
      this.categoryList= res
    }
  })
}
}
