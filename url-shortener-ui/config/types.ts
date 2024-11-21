export interface Config {
    downstream: {
      server: {
        query: {
          host: string;
          port: number;
        };
        command: {
          host: string;
          port: number;
        };
      };
    };
  }