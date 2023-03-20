import { TestBed } from '@angular/core/testing';

import { EnvVariablesServiceService } from './env-variables-service.service';

describe('EnvVariablesServiceService', () => {
  let service: EnvVariablesServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(EnvVariablesServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
