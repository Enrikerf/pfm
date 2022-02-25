
export interface Task{
    uuid : string;
    host : string;
    port : string;
    commands : string[];
    mode : string;
    status : string;
    executionMode : string;
}