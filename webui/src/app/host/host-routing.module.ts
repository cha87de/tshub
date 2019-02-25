import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ListComponent } from './list/list.component';
import { MainComponent } from './main/main.component';

const hostRoutes: Routes = [
    {
        path: 'host',
        component: MainComponent,

        children: [
            {
                path: '',
                component: ListComponent,
                pathMatch: 'full'
            }
        ]
    }
];


@NgModule({
    imports: [
        RouterModule.forChild(hostRoutes)
    ],
    exports: [
        RouterModule
    ]
})
export class HostRoutingModule { }
