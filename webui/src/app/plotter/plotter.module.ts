import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PlotdataComponent } from './plotdata/plotdata.component';

@NgModule({
  declarations: [
    PlotdataComponent
  ],
  imports: [
    CommonModule
  ],
  exports: [
    PlotdataComponent
  ]
})
export class PlotterModule { }
