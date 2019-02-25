import { Component, OnInit, Input, AfterViewInit, ElementRef, ViewChild  } from '@angular/core';
import { DefaultService, Domain, PlotData } from 'src/app/api';
import { Chart } from 'chart.js';

@Component({
	selector: 'plotter-plotdata',
	templateUrl: './plotdata.component.html',
	styleUrls: ['./plotdata.component.css']
})
export class PlotdataComponent implements OnInit, AfterViewInit {

	@Input() domain: Domain;
	@Input() metric: string;

    @ViewChild('canvas') canvas: ElementRef;

	plotdata: PlotData;
	chart: Chart;
	pastData: {}[] = [];
	futureData: {}[] = [];

	constructor(private defaultService: DefaultService) { }

	ngOnInit() {
		this.queryData();
	}

	ngAfterViewInit() {
		this.createChart();
	}

	queryData() {
		const dataQuery = this.defaultService.getDomainPlotdata(this.domain.name, this.metric, "1h")
		dataQuery.toPromise().then((plotdata) => {
			this.plotdata = plotdata;
			this.plotdata.past.forEach((item) => {
				let value = 0;
				let ts = 0;
				if (item.value !== undefined) {
					value = item.value;
				}
				if (item.timestamp !== undefined) {
					ts = item.timestamp;
				}
				let point = {
					x: ts,
					y: value
				};
				// console.info(point, item);
				this.pastData.push(point);
			});

			this.plotdata.past.forEach((item) => {
				let value = 0;
				let ts = 30;
				if (item.value !== undefined) {
					value = item.value;
				}
				if (item.timestamp !== undefined) {
					ts = item.timestamp + 30;
				}
				let point = {
					x: ts,
					y: value
				};
				// console.info(point, item);
				this.futureData.push(point);
			});			
		});
	}

	createChart() {
		const cdata = {
			// labels: Array(this.pastData.length).fill(''),
			datasets: [
				{
					label: 'past',
					borderColor: "#3f51b5",
					data: this.pastData,
					pointRadius: 0.1
				},
				{
					label: 'future',
					borderColor: "#689F38",
					data: this.futureData,
					pointRadius: 0.1
				}				
			]
		};
		const ctx = this.canvas.nativeElement.getContext('2d');
		this.chart = new Chart(ctx, {
			type: 'line',
			data: cdata,
			options: {
				title: {
					display: false
				 },
				 scales: {
					xAxes: [{
					   type: 'linear',
					   /*ticks: {
						  suggestedMin: 0,
						  suggestedMax: 30,
						  stepSize: 1 //interval between ticks
					   }*/
					}],
					yAxes: [{
					   display: true,
					   /*ticks: {
						  suggestedMin: 0,
						  suggestedMax: 100
					   }*/
					}]
				 }				
				/*animation: {
					duration: 0
				},
				responsive: true,
				maintainAspectRatio: false,
				responsiveAnimationDuration: 0,
				legend: {
					display: true,
					labels: {
						boxWidth: 20
					}
				},
				scales: {
					xAxes: [{
						stacked: false,
						display: true,
						ticks: {
							beginAtZero: true,
							steps:1,
							stepValue:1,
							max:30
						}						
					}],
					yAxes: [{
						stacked: false,
						scaleLabel: {
							display: true,
							labelString: 'Utilisation'
						}
					}]
				},
				tooltips: {
				}*/
			}
		});
	}

}
