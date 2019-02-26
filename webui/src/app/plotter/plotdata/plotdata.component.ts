import { Component, OnInit, Input, AfterViewInit, ElementRef, ViewChild } from '@angular/core';
import { DefaultService, Domain, PlotData, Host } from 'src/app/api';
import { Chart } from 'chart.js';
import { SettingsService } from 'src/app/settings.service';
import { Subscription, timer } from 'rxjs';

@Component({
	selector: 'plotter-plotdata',
	templateUrl: './plotdata.component.html',
	styleUrls: ['./plotdata.component.css']
})
export class PlotdataComponent implements OnInit, AfterViewInit {

	@Input() domain: Domain;
	@Input() host: Host;
	@Input() metric: string;

	@ViewChild('canvas') canvas: ElementRef;

	plotdata: PlotData;
	chart: Chart;
	pastData: {}[] = [];
	futureData: {}[] = [];

	reloadSubscription: Subscription;
	autoReloadSubscription: Subscription;

	constructor(private defaultService: DefaultService, private settingsService: SettingsService) { }

	ngOnInit() {
		this.queryData();

		// register to reload event
		this.reloadSubscription = this.settingsService.reloadObservable.subscribe(
			reload => {
				// this.queryData();
			}
		);

		this.autoReload();

	}

	ngOnDestroy() {
		this.reloadSubscription.unsubscribe();
		this.autoReloadSubscription.unsubscribe();
	}

	ngAfterViewInit() {
		this.createChart();
	}

	queryData() {
		const timeframe = this.settingsService.timeframe;
		if (this.domain != undefined) {
			// query domain
			this.queryDataDomain(timeframe);
		} else if (this.host != undefined) {
			// query host
			this.queryDataHost(timeframe);
		}
	}

	autoReload() {
		const source = timer(1000, 2000);
		this.autoReloadSubscription = source.subscribe(val => {
			// reload
			// console.info("reload");
			this.pastData = [];
			this.queryData();
		});
	}

	queryDataDomain(timeframe: string) {
		const dataQuery = this.defaultService.getDomainPlotdata(this.domain.name, this.metric, timeframe)
		dataQuery.toPromise().then((plotdata) => {
			this.handlePlotdata(plotdata);
		});
	}

	queryDataHost(timeframe: string) {
		const dataQuery = this.defaultService.getHostPlotdata(this.host.name, this.metric, timeframe)
		dataQuery.toPromise().then((plotdata) => {
			this.handlePlotdata(plotdata);
		});
	}

	handlePlotdata(plotdata) {
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
		if (this.chart !== undefined) {
			// update chart
			this.chart.data.datasets[0].data = this.pastData;
			this.chart.update();			
		}
	}

	createChart() {
		if (this.chart !== undefined) {
			return;
		}

		const ctx = this.canvas.nativeElement.getContext('2d');
		this.chart = new Chart(ctx, {
			type: 'line',
			data: {
				datasets: [
					{
						label: 'past',
						borderColor: "#3f51b5",
						data: this.pastData,
						pointRadius: 0.1
					}
					/*{
						label: 'future',
						borderColor: "#689F38",
						data: this.futureData,
						pointRadius: 0.1
					}*/
				]
			},
			options: {
				title: {
					display: false
				},
				scales: {
					xAxes: [{
						type: 'linear'
					}],
					yAxes: [{
						display: true,
						ticks: {
							min: 0
						}
					}]
				},
				animation: false
			}
		});

	}

}
