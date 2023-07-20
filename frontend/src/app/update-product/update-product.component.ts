import { Component } from '@angular/core';
import { ProductService } from '../services/product.service';
import { category, product } from '../data-type';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-update-product',
  templateUrl: './update-product.component.html',
  styleUrls: ['./update-product.component.css']
})
export class UpdateProductComponent {
  categoryList:category[] = [];
  product = {} as product
constructor(private productService:ProductService, private route:ActivatedRoute){}

ngOnInit():void{
  let productId = this.route.snapshot.paramMap.get('id')
  productId && this.productService.getProduct(productId).subscribe((data:product)=>{
    this.product = data
  })
  this.productService.categoryList().subscribe((res:category[])=>{
    if(res){
      this.categoryList= res
    }
  })
}
  submitUpdate(data:product){
    this.productService.updateProduct(data)
  }
}
