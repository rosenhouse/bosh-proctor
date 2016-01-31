# proctor
A tool for running BOSH 101 classrooms.

- [Binary releases](https://github.com/rosenhouse/proctor/releases)

- [Pivotal Tracker project](https://www.pivotaltracker.com/n/projects/1434846)

- [Concourse CI Pipeline](https://proctor.ci.cf-app.com/)


### Basic usage
0. Load credentials for your AWS environment
    ```
    export AWS_DEFAULT_REGION=us-east-1  # this can be any region that has a BOSH-lite AMI
    export AWS_ACCESS_KEY_ID=YOUR-ACCESS-KEY
    export AWS_SECRET_ACCESS_KEY=YOUR-SECRET-KEY
    ```
    
0. Create a new classroom
    ```
    proctor create -name my-classroom -number 3
    ```
    This will spin up 3 EC2 instances in your AWS account.
    
0. Watch your classroom get created
    ```
    proctor describe -name my-classroom
    ```
    The SSH key was generated at `create` time and is world-readable.

0. Run a command on all VMs
    ```
    proctor run -name my-classroom -c 'bosh status'
    ```

0. Destroy your classroom
    ```
    proctor destroy -name my-classroom
    ```

