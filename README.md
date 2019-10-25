# 1. Shuffle

Simple CLI application that prints the numbers from 1 to 10 in random order to the terminal.

This application uses an implementation of the "inside-out" variant of the [Fisher-Yates shuffle algorithm](https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle#The_%22inside-out%22_algorithm) to generate a random permutation of the input sequence.

Developed using **Go** version go1.12.6


## Instructions

1. Install go: https://golang.org/doc/install#install
2. Clone this repository into your local machine
3. cd to `shuffle` folder
4. Open a terminal session
5. Build the app running the following command:

		$ go build

7. Run the application running the following command:

		$ ./shuffle


# 2. Application Delivery Controller server

>Imagine a server with the following specs:
>- 4 times Intel(R) Xeon(R) CPU E7-4830 v4 @ 2.00GHz  
>- 64GB of ram  
>- 2 tb HDD disk space  
>- 2 x 10Gbit/s nics

On the proposed scenario, I would consider the following metrics:
- **Uptime rate**: monitor the percentage of time that the server is available in a period of time. Also, performing periodic health checks could be useful to trigger a recovery plan as soon as the server fails.
- **CPU usage**: Given that SSL/TLS is very CPU intensive, it's important to monitor CPU and memory utilization.
- **Temperature of the server processor**: in case of overheating, the server could be shut down automatically to prevent hardware damage 
- **Check SSL certificates for expiration**: to prevent rejected connections due to expired certificates
- **Packet loss rate**: useful to analyze the reliability of the server and/or network
- **Number of terminated connections due to invalid certificates**: useful to detect security attacks
- **Latency and Throughtput**: to detect any network issues that affect connection speed

The approach to collect these metrics would depend on the technology and providers used to deploy the SSL-offloader/Proxy server. For example: 
- If we use a cloud based server, most of the providers (AWS, MS Azure, Google Platform, etc.) would collect these and even include additional metrics to monitor our server.
- If we use a local server, then we should consider using a monitoring solution (e.g. Zabbix, Nagios) that would collect these metrics. Additionally, some of these specialized software would allow us to define our own metrics by creating scripts to collect them and also create triggers to solve or mitigate issues.

I think that the most challenging parts of monitoring this type of servers is the definition of metrics and understanding the collected data. The performance of the SSL-offloader/Proxy server has strong implications on the security of backend applications. This means that weak monitoring may affect the confidentiality, integrity or availability of the apps behind the SSL-offloader/Proxy server.

In conclusion, well defined metrics and a clear understanding of what each metric implies would result in a good visualization of the server's performance, helping the team to quickly identify problems and, therefore, solve or mitigate these issues faster.