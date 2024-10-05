# Installation Instructions

## Using Docker

1. **Rename Configuration File:**
   - Rename `config.yaml.example` to `config.yaml`.

2. **Start Docker Containers:**
   - In the root directory of your project, run:
     ```bash
     docker-compose up -d
     ```
   - Wait until all services are created and started.

3. **Testing:**
   - You can test the application in the frontend by navigating to [http://localhost:3000](http://localhost:3000).
   - Alternatively, use Postman to access the API at: [http://localhost:8888/api/v1/](http://localhost:8888/api/v1/).

4. **Performance Testing:**
   - To evaluate real-time performance with 1000 data entries, use Postman and check the logs.

## Using Traditional Method

1. **Install Dependencies:**
   - Ensure that PostgreSQL and Redis are installed on your localhost. If not, please install them first.

2. **Frontend Setup:**
   - Navigate to the frontend folder:
     ```bash
     cd frontend
     ```
   - Open a terminal and run:
     ```bash
     npm install
     ```
   - After that, start the frontend with:
     ```bash
     npm run dev
     ```

3. **Backend Setup:**
    - find dump.sql in backendfolder then import it:
     ```bash
    psql -U <username> -d <database_name> -f dump.sql
    * For example:
        ```bash
    psql -U postgres -d postgres -f dump.sql
   - Open another terminal and navigate to the backend folder:
     ```bash
     cd backend
     ```
   - Run the following command to tidy up dependencies:
     ```bash
     go mod tidy
     ```

   - Finally, start the backend with:
     ```bash
     go run main.go
     ```
5. **Doocument API**
   - open backend folder:
     ```bash
     and import name file 'bpjstk.postman_collection.json' to your postman