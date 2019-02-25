import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PlotdataComponent } from './plotdata.component';

describe('PlotdataComponent', () => {
  let component: PlotdataComponent;
  let fixture: ComponentFixture<PlotdataComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PlotdataComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PlotdataComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
