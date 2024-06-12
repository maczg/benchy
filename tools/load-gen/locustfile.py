from locust import HttpUser, task, between, constant_pacing
from locust_plugins import constant_total_ips


class User(HttpUser):
    # wait_time = between(1, 2.5)

    @task
    def cpu_load(self):
        self.client.get("/cpu?n=30")

    @task
    def memory_load(self):
         self.client.get("/mem?size=10000000")
