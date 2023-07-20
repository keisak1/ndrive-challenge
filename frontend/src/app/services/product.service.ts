import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { category, product} from '../data-type';
import { map } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ProductService {

  url = "http://localhost:8000/"
  constructor(private http: HttpClient, private router:Router) { }

  addProduct(data:product){

    this.http.post(this.url + 'api/products/create',  data  , {observe: 'response', withCredentials:true} ).subscribe((result)=>{
      if(result){
        this.router.navigate(['home'])
      }
    })
  }


  getProduct(id:string){
    return this.http.get<product>(this.url + `api/products/${id}`).pipe(
      map((response: any) => response.data.result)
    );
  }
  updateProduct(data:product){

    data.price = parseFloat(data.price.toString())
    data.stock = parseInt(data.stock.toString())
    data.rating = parseInt(data.rating.toString())
    return this.http.patch(this.url + `api/products/edit/${data.id}`, data, {observe: 'response', withCredentials:true}).subscribe((result)=>{
      if(result){
        this.router.navigate(['home'])
      }
    })
  }

  categoryList(){
    return this.http.get<category[]>(this.url + 'api/categories').pipe(
      map((response: any) => response.data.categories)
    );
  }

  productList(){
    return this.http.get<product[]>(this.url+'api/products').pipe(
      map((response: any) => response.data.products)
    );
  }
  deleteProduct(id:string){
   return this.http.delete(this.url + `api/products/delete/${id}`)
  }

  searchProduct(data:string){
    return this.http.get<product[]>(this.url + `api/products?search=${data.search}`).pipe(
      map((response: any) => response.data.products)
    );
  }
}
