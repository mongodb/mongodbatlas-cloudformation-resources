import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
  <div class="container-md">
    <router-outlet></router-outlet>
  </div>
  `
})
export class AppComponent { }
