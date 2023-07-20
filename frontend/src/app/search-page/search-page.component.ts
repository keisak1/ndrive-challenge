import { Component } from '@angular/core';
import { product } from '../data-type';
import { ProductService } from '../services/product.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-search-page',
  templateUrl: './search-page.component.html',
  styleUrls: ['./search-page.component.css']
})
export class SearchPageComponent {
  productList:product[] = [];
  menuType:String = 'default'
  constructor(private productService:ProductService, private router:Router, private route:ActivatedRoute) {}
  ngOnInit(): void {
    let data = this.route.snapshot.paramMap.get('query')

    data && this.productService.searchProduct(data).subscribe((res)=>{
      this.productList = res
    })
  }
  search(data:string){
    console.warn(data)
    this.productService.searchProduct(data).subscribe((res)=>{
      this.productList = res
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
