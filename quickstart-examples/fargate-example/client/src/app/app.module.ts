import { APP_INITIALIZER, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { EmployeesListComponent } from './employees-list/employees-list.component';
import { EmployeeFormComponent } from './employee-form/employee-form.component';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { AddEmployeeComponent } from './add-employee/add-employee.component';
import { EditEmployeeComponent } from './edit-employee/edit-employee.component'; // <-- add this line
import { EnvVariablesServiceService } from './env-variables-service.service';

export function envVariableLoader(
  envVariableService: EnvVariablesServiceService
) {
  return () => {
    return envVariableService.getEnv();
  };
}
@NgModule({
  declarations: [
    AppComponent,
    EmployeesListComponent,
    EmployeeFormComponent,
    AddEmployeeComponent,
    EditEmployeeComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule // <-- add this line
  ],
  providers: [{
                    provide: APP_INITIALIZER,
                    useFactory: envVariableLoader,
                    deps: [EnvVariablesServiceService],
                    multi: true,
                  },],
  bootstrap: [AppComponent]
})
export class AppModule { }
