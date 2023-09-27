from locust import HttpUser, task, between

class User(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def cpu_load(self):
        self.client.get("/cpu?n=30")

    @task
    def memory_load(self):
         self.client.get("/mem?size=10000000")
