import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WindChart } from './wind-chart';

describe('WindChart', () => {
  let component: WindChart;
  let fixture: ComponentFixture<WindChart>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WindChart]
    })
    .compileComponents();

    fixture = TestBed.createComponent(WindChart);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
