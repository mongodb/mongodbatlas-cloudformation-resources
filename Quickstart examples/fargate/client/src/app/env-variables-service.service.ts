import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { lastValueFrom } from 'rxjs';
@Injectable({
  providedIn: 'root',
})
export class EnvVariablesServiceService {
  public ENV: any;
  constructor(private http: HttpClient) {}

  async getEnv() {
    try {
      console.log('Loading Env....')
      this.ENV = await lastValueFrom(this.http.get('getEnv'));
      console.log('Loaded ENV',this.ENV)
    } catch (e) {
      console.log('Loading Env Failed.....')
      console.log(e)
      return null
    }

    return this.ENV;
  }
}
