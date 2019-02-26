import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListComponent } from './list/list.component';
import { MainComponent } from './main/main.component';
import { HostRoutingModule } from './host-routing.module';

import { HttpClientModule } from '@angular/common/http';
import { MatButtonModule } from '@angular/material';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatSelectModule } from '@angular/material/select';
import { ApiModule } from '../api';
import { PipesModule } from '../pipes/pipes.module';
import { PlotterModule } from '../plotter/plotter.module';

@NgModule({
  declarations: [ ListComponent, MainComponent ],
  imports: [
    CommonModule,
    HostRoutingModule,
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
export class HostModule { }
