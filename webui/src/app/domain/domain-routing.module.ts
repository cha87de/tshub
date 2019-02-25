import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ListComponent } from './list/list.component';
import { MainComponent } from './main/main.component';

const domainRoutes: Routes = [
    {
        path: 'domain',
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
        RouterModule.forChild(domainRoutes)
    ],
    exports: [
        RouterModule
    ]
})
export class DomainRoutingModule { }
