import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListComponent } from './list/list.component';
import { DomainRoutingModule } from './domain-routing.module';

import { MatButtonModule } from '@angular/material';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatSelectModule } from '@angular/material/select';
import { MainComponent } from './main/main.component';

import { HttpClientModule } from '@angular/common/http';
import { ApiModule } from '../api';
import { PipesModule } from '../pipes/pipes.module';
import { PlotterModule } from '../plotter/plotter.module';

@NgModule({
  declarations: [MainComponent, ListComponent],
  imports: [
    DomainRoutingModule,
    CommonModule,
    HttpClientModule,
    ApiModule,
    PipesModule,
    PlotterModule,

    // material modules
    MatButtonModule,
    MatToolbarModule,
    MatSidenavModule,
    MatCardModule,
    MatListModule,
    MatGridListModule,
    MatSelectModule        
  ]
})
export class DomainModule { }
