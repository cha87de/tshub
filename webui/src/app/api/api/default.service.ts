/**
 * TSHub API
 * The REST API for the TimeSeries-Profiler
 *
 * OpenAPI spec version: 0.1
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */
/* tslint:disable:no-unused-variable member-ordering */

import { Inject, Injectable, Optional }                      from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams,
         HttpResponse, HttpEvent }                           from '@angular/common/http';
import { CustomHttpUrlEncodingCodec }                        from '../encoder';

import { Observable }                                        from 'rxjs/Observable';

import { AnyValue } from '../model/anyValue';
import { Domain } from '../model/domain';
import { DomainDetails } from '../model/domainDetails';
import { Host } from '../model/host';
import { HostDetails } from '../model/hostDetails';
import { PlotData } from '../model/plotData';

import { BASE_PATH, COLLECTION_FORMATS }                     from '../variables';
import { Configuration }                                     from '../configuration';


@Injectable()
export class DefaultService {

    protected basePath = 'http://localhost:8080/v0.1';
    public defaultHeaders = new HttpHeaders();
    public configuration = new Configuration();

    constructor(protected httpClient: HttpClient, @Optional()@Inject(BASE_PATH) basePath: string, @Optional() configuration: Configuration) {
        if (basePath) {
            this.basePath = basePath;
        }
        if (configuration) {
            this.configuration = configuration;
            this.basePath = basePath || configuration.basePath || this.basePath;
        }
    }

    /**
     * @param consumes string[] mime-types
     * @return true: consumes contains 'multipart/form-data', false: otherwise
     */
    private canConsumeForm(consumes: string[]): boolean {
        const form = 'multipart/form-data';
        for (let consume of consumes) {
            if (form === consume) {
                return true;
            }
        }
        return false;
    }


    /**
     * Get details of given domain by its name
     * 
     * @param domainname 
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getDomain(domainname: string, observe?: 'body', reportProgress?: boolean): Observable<DomainDetails>;
    public getDomain(domainname: string, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<DomainDetails>>;
    public getDomain(domainname: string, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<DomainDetails>>;
    public getDomain(domainname: string, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {
        if (domainname === null || domainname === undefined) {
            throw new Error('Required parameter domainname was null or undefined when calling getDomain.');
        }

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        let httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set("Accept", httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        let consumes: string[] = [
        ];

        return this.httpClient.get<DomainDetails>(`${this.basePath}/domain/${encodeURIComponent(String(domainname))}`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns the plotdata (past and prediction) of given domain and metric
     * 
     * @param domainname 
     * @param metric 
     * @param timeframe 
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getDomainPlotdata(domainname: string, metric: string, timeframe?: string, observe?: 'body', reportProgress?: boolean): Observable<PlotData>;
    public getDomainPlotdata(domainname: string, metric: string, timeframe?: string, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<PlotData>>;
    public getDomainPlotdata(domainname: string, metric: string, timeframe?: string, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<PlotData>>;
    public getDomainPlotdata(domainname: string, metric: string, timeframe?: string, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {
        if (domainname === null || domainname === undefined) {
            throw new Error('Required parameter domainname was null or undefined when calling getDomainPlotdata.');
        }
        if (metric === null || metric === undefined) {
            throw new Error('Required parameter metric was null or undefined when calling getDomainPlotdata.');
        }

        let queryParameters = new HttpParams({encoder: new CustomHttpUrlEncodingCodec()});
        if (metric !== undefined) {
            queryParameters = queryParameters.set('metric', <any>metric);
        }
        if (timeframe !== undefined) {
            queryParameters = queryParameters.set('timeframe', <any>timeframe);
        }

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        let httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set("Accept", httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        let consumes: string[] = [
        ];

        return this.httpClient.get<PlotData>(`${this.basePath}/domain/${encodeURIComponent(String(domainname))}/plotdata`,
            {
                params: queryParameters,
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Get a list of available domains (virtual machines / containers)
     * 
     * @param hostname 
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getDomains(hostname?: string, observe?: 'body', reportProgress?: boolean): Observable<Array<Domain>>;
    public getDomains(hostname?: string, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Array<Domain>>>;
    public getDomains(hostname?: string, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Array<Domain>>>;
    public getDomains(hostname?: string, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        let queryParameters = new HttpParams({encoder: new CustomHttpUrlEncodingCodec()});
        if (hostname !== undefined) {
            queryParameters = queryParameters.set('hostname', <any>hostname);
        }

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        let httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set("Accept", httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        let consumes: string[] = [
        ];

        return this.httpClient.get<Array<Domain>>(`${this.basePath}/domains`,
            {
                params: queryParameters,
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Get details of the given host
     * 
     * @param hostname 
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getHost(hostname: string, observe?: 'body', reportProgress?: boolean): Observable<HostDetails>;
    public getHost(hostname: string, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<HostDetails>>;
    public getHost(hostname: string, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<HostDetails>>;
    public getHost(hostname: string, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {
        if (hostname === null || hostname === undefined) {
            throw new Error('Required parameter hostname was null or undefined when calling getHost.');
        }

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        let httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set("Accept", httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        let consumes: string[] = [
        ];

        return this.httpClient.get<HostDetails>(`${this.basePath}/host/${encodeURIComponent(String(hostname))}`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Returns the plotdata (past and prediction) of given host and metric
     * 
     * @param hostname 
     * @param metric 
     * @param timeframe 
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getHostPlotdata(hostname: string, metric: string, timeframe?: string, observe?: 'body', reportProgress?: boolean): Observable<PlotData>;
    public getHostPlotdata(hostname: string, metric: string, timeframe?: string, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<PlotData>>;
    public getHostPlotdata(hostname: string, metric: string, timeframe?: string, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<PlotData>>;
    public getHostPlotdata(hostname: string, metric: string, timeframe?: string, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {
        if (hostname === null || hostname === undefined) {
            throw new Error('Required parameter hostname was null or undefined when calling getHostPlotdata.');
        }
        if (metric === null || metric === undefined) {
            throw new Error('Required parameter metric was null or undefined when calling getHostPlotdata.');
        }

        let queryParameters = new HttpParams({encoder: new CustomHttpUrlEncodingCodec()});
        if (metric !== undefined) {
            queryParameters = queryParameters.set('metric', <any>metric);
        }
        if (timeframe !== undefined) {
            queryParameters = queryParameters.set('timeframe', <any>timeframe);
        }

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        let httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set("Accept", httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        let consumes: string[] = [
        ];

        return this.httpClient.get<PlotData>(`${this.basePath}/host/${encodeURIComponent(String(hostname))}/plotdata`,
            {
                params: queryParameters,
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Get a list of available physical hosts
     * 
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getHosts(observe?: 'body', reportProgress?: boolean): Observable<Array<Host>>;
    public getHosts(observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Array<Host>>>;
    public getHosts(observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Array<Host>>>;
    public getHosts(observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        let httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set("Accept", httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        let consumes: string[] = [
        ];

        return this.httpClient.get<Array<Host>>(`${this.basePath}/hosts`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Get the full profile by given name
     * 
     * @param profilename 
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getProfile(profilename: string, observe?: 'body', reportProgress?: boolean): Observable<AnyValue>;
    public getProfile(profilename: string, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<AnyValue>>;
    public getProfile(profilename: string, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<AnyValue>>;
    public getProfile(profilename: string, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {
        if (profilename === null || profilename === undefined) {
            throw new Error('Required parameter profilename was null or undefined when calling getProfile.');
        }

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        let httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set("Accept", httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        let consumes: string[] = [
        ];

        return this.httpClient.get<AnyValue>(`${this.basePath}/profile/${encodeURIComponent(String(profilename))}`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Get a list of names from stored profiles
     * 
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public getProfileNames(observe?: 'body', reportProgress?: boolean): Observable<Array<string>>;
    public getProfileNames(observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<Array<string>>>;
    public getProfileNames(observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<Array<string>>>;
    public getProfileNames(observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        let httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set("Accept", httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        let consumes: string[] = [
        ];

        return this.httpClient.get<Array<string>>(`${this.basePath}/profiles`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

}
