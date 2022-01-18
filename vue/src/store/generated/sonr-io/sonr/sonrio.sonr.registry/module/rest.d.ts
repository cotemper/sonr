export interface ProtobufAny {
    "@type"?: string;
}
/**
* Did represents a string that has been parsed and validated as a DID. The parts are stored
in the individual fields.
*/
export interface RegistryDid {
    /** Method is the method used to create the DID. For the Sonr network it is "sonr". */
    method?: string;
    /** Network is the network the DID is on. For testnet it is "testnet". i.e "did:sonr:testnet:". */
    network?: string;
    id?: string;
    /** Paths is a list of paths that the DID is valid for. This is used to identify the Service. */
    paths?: string[];
    /** Query is the query string that was used to create the DID. This is followed by a '?'. */
    query?: string;
    /** Fragment is the fragment string that was used to create the DID. This is followed by a '#'. */
    fragment?: string;
}
export interface RegistryMsgCreateAccountResponse {
    /** @format int32 */
    code?: number;
    /**
     * Did represents a string that has been parsed and validated as a DID. The parts are stored
     * in the individual fields.
     */
    did?: RegistryDid;
}
export interface RegistryMsgRegisterNameResponse {
    /** @format int32 */
    code?: number;
    /**
     * Did represents a string that has been parsed and validated as a DID. The parts are stored
     * in the individual fields.
     */
    did?: RegistryDid;
}
export interface RegistryMsgRegisterServiceResponse {
    /** @format int32 */
    code?: number;
    /**
     * Did represents a string that has been parsed and validated as a DID. The parts are stored
     * in the individual fields.
     */
    did?: RegistryDid;
}
/**
 * Params defines the parameters for the module.
 */
export declare type RegistryParams = object;
/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface RegistryQueryParamsResponse {
    /** params holds all the parameters of this module. */
    params?: RegistryParams;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title registry/did.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/sonrio/sonr/registry/params
     */
    queryParams: (params?: RequestParams) => Promise<HttpResponse<RegistryQueryParamsResponse, RpcStatus>>;
}
export {};
