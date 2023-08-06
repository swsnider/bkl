# bkl

Layered configuration language

![Go Reference](data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iOTAiIGhlaWdodD0iMjAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PGcgZmlsbD0ibm9uZSIgZmlsbC1ydWxlPSJldmVub2RkIj48cGF0aCBkPSJNMiAwaDI2djIwSDJhMiAyIDAgMDEtMi0yVjJhMiAyIDAgMDEyLTJ6IiBmaWxsPSIjNUM1QzVDIi8+PHBhdGggZD0iTTg3Ljk5IDBIMjh2MjBoNTkuOTlhMiAyIDAgMDAyLTJWMmEyIDIgMCAwMC0yLTJ6IiBmaWxsPSIjMDA3RDlDIi8+PHBhdGggZD0iTTM1LjE3NyAxNHYtMi45MTVjMC0uMzc0LjA3Mi0uNjgyLjIxNS0uOTI0LjE0Mi0uMjQyLjMyNC0uNDIzLjU0NC0uNTQ0LjIyLS4xMjIuNDQ0LS4xODIuNjcxLS4xODIuMTY5IDAgLjMwMy4wMS40MDEuMDI3LjEuMDE5LjE5LjA0My4yNy4wNzJsLjE2NS0xLjAzNGEuOTc3Ljk3NyAwIDAwLS4yNzUtLjA4MiAyLjAwNCAyLjAwNCAwIDAwLS4zMy0uMDI4Yy0uNDI1IDAtLjc4My4wOTMtMS4wNzIuMjgtLjI5LjE4OC0uNTA4LjQ2LS42NTUuODJsLS4xMS0uOTlIMzRWMTRoMS4xNzd6bTUuNjYuMTFjLjU0MiAwIDEuMDE5LS4xMTQgMS40My0uMzQxLjQxLS4yMjcuNzE0LS41NDMuOTEyLS45NDZsLS45MzUtLjQ0Yy0uMTQuMjU3LS4zMjIuNDUzLS41NS41ODktLjIyNy4xMzUtLjUyLjIwMy0uODguMjAzLS4zMyAwLS42MTYtLjA3My0uODU4LS4yMmExLjQzOSAxLjQzOSAwIDAxLS41Ni0uNjQ5IDIuMTY5IDIuMTY5IDAgMDEtLjE3Ny0uNjU1bC0uMDA0LS4wNDloNC4wM2MuMDEtLjA1OS4wMTgtLjEyNC4wMjUtLjE5NmwuMDA4LS4xMTJjLjAwOC0uMTE3LjAxMS0uMjQyLjAxMS0uMzc0IDAtLjQ4NC0uMS0uOTE4LS4zMDItMS4zMDNhMi4xOSAyLjE5IDAgMDAtLjg3LS45MDNjLS4zNzctLjIxNi0uODI2LS4zMjQtMS4zNDctLjMyNC0uNTI4IDAtLjk5NS4xMTUtMS40MDIuMzQ2YTIuNDQyIDIuNDQyIDAgMDAtLjk1Ny45OWMtLjIzMS40My0uMzQ3LjkzNy0uMzQ3IDEuNTI0cy4xMTYgMS4wOTUuMzQ3IDEuNTI0Yy4yMy40MjkuNTU1Ljc1OS45NzMuOTkuNDE4LjIzLjkwMi4zNDYgMS40NTIuMzQ2em0xLjMwMi0zLjQxaC0yLjg5NWwuMDItLjExNWMuMDE2LS4wNjguMDMzLS4xMzQuMDUzLS4xOTdsLjA2Ny0uMTgzYy4xMjUtLjI5My4zMDMtLjUxMy41MzQtLjY2LjIzLS4xNDcuNTA4LS4yMi44My0uMjIuNDcgMCAuODI5LjE1NiAxLjA3OC40NjguMTU2LjE5NC4yNTUuNDM3LjI5Ni43MjhsLjAxNy4xNzl6TTQ2LjMyIDE0VjkuNDI0aDEuNzVWOC41aC0xLjc4M3YtLjcwNGMwLS4yNy4wNjQtLjQ3Ni4xOS0uNjE3bC4wNjktLjA2NWMuMTcyLS4xNC40MDktLjIwOS43MS0uMjA5LjEzMSAwIC4yNDkuMDEzLjM1MS4wMzhhLjkyNy45MjcgMCAwMS4yNzUuMTE2bC4yMi0uODQ3YTEuMDYgMS4wNiAwIDAwLS40MjMtLjE5MiAyLjUxNiAyLjUxNiAwIDAwLS42MjItLjA3MmMtLjI5MyAwLS41NTkuMDM4LS43OTcuMTE1YTEuNjggMS42OCAwIDAwLS42MS4zNDFjLS4xNy4xNS0uMjk2LjMzNi0uMzguNTU2LS4wODUuMjItLjEyNy40Ny0uMTI3Ljc0OFY4LjVoLTEuMDM0di45MjRoMS4wMzRWMTRoMS4xNzd6bTUuMS4xMWMuNTQzIDAgMS4wMi0uMTE0IDEuNDMtLjM0MS40MTEtLjIyNy43MTUtLjU0My45MTMtLjk0NmwtLjkzNS0uNDRjLS4xMzkuMjU3LS4zMjIuNDUzLS41NS41ODktLjIyNy4xMzUtLjUyLjIwMy0uODguMjAzLS4zMyAwLS42MTYtLjA3My0uODU4LS4yMmExLjQzOSAxLjQzOSAwIDAxLS41Ni0uNjQ5IDIuMTY5IDIuMTY5IDAgMDEtLjE3Ny0uNjU1bC0uMDA0LS4wNDloNC4wM2MuMDEtLjA1OS4wMTgtLjEyNC4wMjUtLjE5NmwuMDA4LS4xMTJjLjAwOC0uMTE3LjAxMS0uMjQyLjAxMS0uMzc0IDAtLjQ4NC0uMS0uOTE4LS4zMDItMS4zMDNhMi4xOSAyLjE5IDAgMDAtLjg2OS0uOTAzYy0uMzc4LS4yMTYtLjgyNy0uMzI0LTEuMzQ4LS4zMjQtLjUyOCAwLS45OTUuMTE1LTEuNDAyLjM0NmEyLjQ0MiAyLjQ0MiAwIDAwLS45NTcuOTljLS4yMzEuNDMtLjM0Ny45MzctLjM0NyAxLjUyNHMuMTE2IDEuMDk1LjM0NyAxLjUyNGMuMjMxLjQyOS41NTUuNzU5Ljk3My45OS40MTguMjMuOTAyLjM0NiAxLjQ1Mi4zNDZ6bTEuMzAzLTMuNDFoLTIuODk1bC4wMi0uMTE1Yy4wMTYtLjA2OC4wMzMtLjEzNC4wNTMtLjE5N2wuMDY3LS4xODNjLjEyNS0uMjkzLjMwMy0uNTEzLjUzNC0uNjYuMjMxLS4xNDcuNTA4LS4yMi44My0uMjIuNDcgMCAuODMuMTU2IDEuMDc4LjQ2OC4xNTYuMTk0LjI1NS40MzcuMjk2LjcyOGwuMDE3LjE3OXpNNTYuNTMgMTR2LTIuOTE1YzAtLjM3NC4wNzItLjY4Mi4yMTUtLjkyNC4xNDMtLjI0Mi4zMjQtLjQyMy41NDQtLjU0NC4yMi0uMTIyLjQ0NC0uMTgyLjY3MS0uMTgyLjE3IDAgLjMwMy4wMS40MDIuMDI3LjA5OS4wMTkuMTg5LjA0My4yNy4wNzJsLjE2NC0xLjAzNGEuOTc3Ljk3NyAwIDAwLS4yNzUtLjA4MiAyLjAwNCAyLjAwNCAwIDAwLS4zMy0uMDI4Yy0uNDI1IDAtLjc4Mi4wOTMtMS4wNzIuMjgtLjI5LjE4OC0uNTA4LjQ2LS42NTUuODJsLS4xMS0uOTloLTFWMTRoMS4xNzZ6bTUuNjYuMTFjLjU0MiAwIDEuMDE5LS4xMTQgMS40My0uMzQxLjQxLS4yMjcuNzE1LS41NDMuOTEzLS45NDZsLS45MzUtLjQ0Yy0uMTQuMjU3LS4zMjMuNDUzLS41NS41ODktLjIyOC4xMzUtLjUyMS4yMDMtLjg4LjIwMy0uMzMgMC0uNjE2LS4wNzMtLjg1OC0uMjJhMS40MzkgMS40MzkgMCAwMS0uNTYxLS42NDkgMi4xNjkgMi4xNjkgMCAwMS0uMTc2LS42NTVsLS4wMDUtLjA0OWg0LjAzYy4wMS0uMDU5LjAxOS0uMTI0LjAyNS0uMTk2bC4wMDktLjExMmMuMDA3LS4xMTcuMDEtLjI0Mi4wMS0uMzc0IDAtLjQ4NC0uMS0uOTE4LS4zMDItMS4zMDNhMi4xOSAyLjE5IDAgMDAtLjg2OS0uOTAzYy0uMzc3LS4yMTYtLjgyNy0uMzI0LTEuMzQ3LS4zMjQtLjUyOCAwLS45OTYuMTE1LTEuNDAzLjM0NmEyLjQ0MiAyLjQ0MiAwIDAwLS45NTcuOTljLS4yMy40My0uMzQ2LjkzNy0uMzQ2IDEuNTI0cy4xMTUgMS4wOTUuMzQ2IDEuNTI0Yy4yMzEuNDI5LjU1Ni43NTkuOTc0Ljk5LjQxOC4yMy45MDIuMzQ2IDEuNDUyLjM0NnptMS4zMDMtMy40MWgtMi44OTZsLjAyMS0uMTE1Yy4wMTUtLjA2OC4wMzItLjEzNC4wNTItLjE5N2wuMDY4LS4xODNjLjEyNC0uMjkzLjMwMi0uNTEzLjUzMy0uNjYuMjMxLS4xNDcuNTA4LS4yMi44My0uMjIuNDcgMCAuODMuMTU2IDEuMDc5LjQ2OC4xNTYuMTk0LjI1NC40MzcuMjk1LjcyOGwuMDE4LjE3OXpNNjcuMyAxNHYtMi45MjZjMC0uNDE4LjA3NS0uNzUyLjIyNS0xLjAwMS4xNS0uMjUuMzQtLjQzLjU2Ny0uNTQ1LjIyNy0uMTEzLjQ2Mi0uMTcuNzA0LS4xNy4zMzcgMCAuNjE3LjEwNC44NDEuMzEzLjIyNC4yMS4zMzYuNTg1LjMzNiAxLjEyOFYxNGgxLjE3N3YtMy41NjRjMC0uNDYyLS4wOTItLjg0My0uMjc1LTEuMTQ0YTEuNzQgMS43NCAwIDAwLS43NDMtLjY3NiAyLjM4NyAyLjM4NyAwIDAwLTEuMDUtLjIyNmMtLjMwMSAwLS41ODUuMDQ4LS44NTMuMTQzYTEuNzIyIDEuNzIyIDAgMDAtLjY5My40NTdjLS4xMTYuMTI1LS4yMTcuMjctLjMuNDMzbC0uMDE3LjAzNy0uMDUxLS45NmgtMS4wNDVWMTRINjcuM3ptOC4wNDMuMTFjLjQxOCAwIC43ODgtLjA2OCAxLjExLS4yMDRhMi4xOSAyLjE5IDAgMDAuODEtLjU3MmMuMjE1LS4yNDUuMzc1LS41MzMuNDc4LS44NjNsLTEuMDY3LS4zNjNjLS4wNDQuMjI3LS4xMjcuNDItLjI0OC41NzdhMS4xOTcgMS4xOTcgMCAwMS0uNDU2LjM2NCAxLjU0IDEuNTQgMCAwMS0uNjUuMTI2Yy0uMzIyIDAtLjYtLjA3My0uODM1LS4yMmExLjM5NyAxLjM5NyAwIDAxLS41NC0uNjQ5Yy0uMTI0LS4yODYtLjE4Ni0uNjM0LS4xODYtMS4wNDUgMC0uNDE4LjA2LS43Ny4xODEtMS4wNTYuMTIxLS4yODYuMjk5LS41MDQuNTM0LS42NTUuMjM0LS4xNS41MTctLjIyNS44NDctLjIyNS4zMzcgMCAuNjEuMDkuODIuMjcuMjA4LjE4LjM2LjQzOC40NTYuNzc1bDEuMTEtLjQ0Yy0uMTEtLjMtLjI3LS41NjgtLjQ4My0uODAzYTIuMTQ4IDIuMTQ4IDAgMDAtLjc4Ny0uNTQ1Yy0uMzEyLS4xMjgtLjY4Ny0uMTkyLTEuMTI3LS4xOTItLjUyOCAwLTEgLjExNC0xLjQxNC4zNDFhMi4zODYgMi4zODYgMCAwMC0uOTY4Ljk4NWMtLjIzLjQyOS0uMzQ2Ljk0LS4zNDYgMS41MzQgMCAuNTk0LjExNSAxLjEwNS4zNDYgMS41MzQuMjMxLjQzLjU1Ni43NTguOTc0Ljk4NS40MTguMjI3Ljg5OC4zNDEgMS40NC4zNDF6bTYuMTkyIDBjLjU0MyAwIDEuMDItLjExNCAxLjQzLS4zNDEuNDExLS4yMjcuNzE1LS41NDMuOTEzLS45NDZsLS45MzUtLjQ0Yy0uMTM5LjI1Ny0uMzIyLjQ1My0uNTUuNTg5LS4yMjcuMTM1LS41Mi4yMDMtLjg4LjIwMy0uMzMgMC0uNjE2LS4wNzMtLjg1OC0uMjJhMS40MzkgMS40MzkgMCAwMS0uNTYtLjY0OSAyLjE2OSAyLjE2OSAwIDAxLS4xNzctLjY1NWwtLjAwNC0uMDQ5aDQuMDNjLjAxLS4wNTkuMDE4LS4xMjQuMDI1LS4xOTZsLjAwOC0uMTEyYy4wMDgtLjExNy4wMTEtLjI0Mi4wMTEtLjM3NCAwLS40ODQtLjEtLjkxOC0uMzAyLTEuMzAzYTIuMTkgMi4xOSAwIDAwLS44Ny0uOTAzYy0uMzc3LS4yMTYtLjgyNi0uMzI0LTEuMzQ3LS4zMjQtLjUyOCAwLS45OTUuMTE1LTEuNDAyLjM0NmEyLjQ0MiAyLjQ0MiAwIDAwLS45NTcuOTljLS4yMzEuNDMtLjM0Ny45MzctLjM0NyAxLjUyNHMuMTE2IDEuMDk1LjM0NyAxLjUyNGMuMjMuNDI5LjU1NS43NTkuOTczLjk5LjQxOC4yMy45MDIuMzQ2IDEuNDUyLjM0NnptMS4zMDMtMy40MWgtMi44OTVsLjAyLS4xMTVjLjAxNi0uMDY4LjAzMy0uMTM0LjA1My0uMTk3bC4wNjctLjE4M2MuMTI1LS4yOTMuMzAzLS41MTMuNTM0LS42Ni4yMy0uMTQ3LjUwOC0uMjIuODMtLjIyLjQ3IDAgLjgzLjE1NiAxLjA3OC40NjguMTU2LjE5NC4yNTUuNDM3LjI5Ni43MjhsLjAxNy4xNzl6TTYuMzU4IDguNzM2Yy0uMDM1IDAtLjA0My0uMDE3LS4wMjYtLjA0NGwuMTg0LS4yMzZjLjAxOC0uMDI2LjA2MS0uMDQ0LjA5Ni0uMDQ0aDMuMTNjLjAzNCAwIC4wNDMuMDI2LjAyNi4wNTNsLS4xNS4yMjdjLS4wMTcuMDI3LS4wNjEuMDUzLS4wODcuMDUzbC0zLjE3My0uMDA5em0tMS4zMjMuODA2Yy0uMDM1IDAtLjA0NC0uMDE3LS4wMjYtLjA0M2wuMTg0LS4yMzdjLjAxNy0uMDI2LjA2MS0uMDQzLjA5Ni0uMDQzaDMuOTk2Yy4wMzYgMCAuMDUzLjAyNi4wNDQuMDUybC0uMDcuMjFjLS4wMDkuMDM1LS4wNDMuMDUzLS4wNzkuMDUzbC00LjE0NS4wMDl6bTIuMTIxLjgwN2MtLjAzNSAwLS4wNDQtLjAyNi0uMDI2LS4wNTNsLjEyMi0uMjE5Yy4wMTgtLjAyNi4wNTMtLjA1Mi4wODgtLjA1MmgxLjc1M2MuMDM1IDAgLjA1Mi4wMjYuMDUyLjA2MWwtLjAxNy4yMWMwIC4wMzYtLjAzNS4wNjItLjA2Mi4wNjJsLTEuOTEtLjAxem05LjA5Ny0xLjc3Yy0uNTUyLjE0LS45My4yNDUtMS40NzIuMzg1LS4xMzIuMDM1LS4xNC4wNDQtLjI1NS0uMDg3LS4xMy0uMTUtLjIyNy0uMjQ2LS40MTEtLjMzMy0uNTUzLS4yNzItMS4wODgtLjE5My0xLjU4Ny4xMy0uNTk2LjM4Ni0uOTAyLjk1Ni0uODkzIDEuNjY2LjAwOC43MDEuNDkgMS4yOCAxLjE4MiAxLjM3Ni41OTcuMDc5IDEuMDk1LS4xMzEgMS40OS0uNTc4LjA4LS4wOTcuMTQ5LS4yMDIuMjM3LS4zMjVoLTEuNjkxYy0uMTg0IDAtLjIyOC0uMTE0LS4xNjctLjI2My4xMTQtLjI3MS4zMjQtLjcyNy40NDctLjk1NWEuMjM2LjIzNiAwIDAxLjIxOS0uMTRoMy4xOWMtLjAxNy4yMzctLjAxNy40NzMtLjA1My43MWEzLjczNyAzLjczNyAwIDAxLS43MTggMS43MThjLS42MzEuODMyLTEuNDU1IDEuMzUtMi40OTggMS40ODktLjg1OC4xMTQtMS42NTYtLjA1Mi0yLjM1Ny0uNTc3YTIuNzUzIDIuNzUzIDAgMDEtMS4xMTMtMS45NDdjLS4xMTQtLjk1NS4xNjctMS44MTQuNzQ1LTIuNTY4LjYyMi0uODE0IDEuNDQ2LTEuMzMyIDIuNDU0LTEuNTE1LjgyMy0uMTUgMS42MTItLjA1MyAyLjMyMi40MjkuNDY1LjMwNy43OTcuNzI3IDEuMDE3IDEuMjM2LjA1Mi4wNzguMDE3LjEyMy0uMDg4LjE0OW0yLjkwMSA0Ljg0NmMtLjc5OC0uMDE4LTEuNTI1LS4yNDYtMi4xMzktLjc3MWEyLjc0OSAyLjc0OSAwIDAxLS45NDYtMS42OTJjLS4xNTgtLjk5LjExNC0xLjg2Ni43MS0yLjY0Ni42NC0uODQyIDEuNDEtMS4yOCAyLjQ1NC0xLjQ2NC44OTQtLjE1NyAxLjczNS0uMDcgMi40OTcuNDQ3LjY5My40NzMgMS4xMjIgMS4xMTMgMS4yMzYgMS45NTQuMTQ5IDEuMTgzLS4xOTIgMi4xNDctMS4wMDggMi45NzFhNC4wMTQgNC4wMTQgMCAwMS0yLjEwMyAxLjEyMmMtLjIzNy4wNDQtLjQ3My4wNTItLjcwMS4wOHptMi4wODUtMy41NGMtLjAwOC0uMTE0LS4wMDgtLjIwMi0uMDI1LS4yOS0uMTU4LS44NjgtLjk1Ni0xLjM1OC0xLjc4OC0xLjE2NS0uODE2LjE4My0xLjM0MS43LTEuNTM0IDEuNTI0YTEuNDQ0IDEuNDQ0IDAgMDAuODA2IDEuNjU3Yy40ODIuMjEuOTY0LjE4NCAxLjQyOS0uMDUyLjY5Mi0uMzYgMS4wNjktLjkyIDEuMTEzLTEuNjc0aC0uMDAxeiIgZmlsbD0iI0ZBRkFGQSIgZmlsbC1ydWxlPSJub256ZXJvIi8+PC9nPjwvc3ZnPg==)
![GitHub: bkl](data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB3aWR0aD0iOTEiIGhlaWdodD0iMjAiIHJvbGU9ImltZyIgYXJpYS1sYWJlbD0iR2l0SHViOiBia2wiPjx0aXRsZT5HaXRIdWI6IGJrbDwvdGl0bGU+PGNsaXBQYXRoIGlkPSJyIj48cmVjdCB3aWR0aD0iOTEiIGhlaWdodD0iMjAiIHJ4PSIzIiBmaWxsPSIjZmZmIi8+PC9jbGlwUGF0aD48ZyBjbGlwLXBhdGg9InVybCgjcikiPjxyZWN0IHdpZHRoPSI2NCIgaGVpZ2h0PSIyMCIgZmlsbD0iIzVjNWM1YyIvPjxyZWN0IHg9IjY0IiB3aWR0aD0iMjciIGhlaWdodD0iMjAiIGZpbGw9IiMwMjdkOWMiLz48L2c+PGcgZmlsbD0iI2ZmZiIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZm9udC1mYW1pbHk9IlZlcmRhbmEsR2VuZXZhLERlamFWdSBTYW5zLHNhbnMtc2VyaWYiIHRleHQtcmVuZGVyaW5nPSJnZW9tZXRyaWNQcmVjaXNpb24iIGZvbnQtc2l6ZT0iMTEwIj48aW1hZ2UgeD0iNSIgeT0iMyIgd2lkdGg9IjE0IiBoZWlnaHQ9IjE0IiB4bGluazpocmVmPSJkYXRhOmltYWdlL3N2Zyt4bWw7YmFzZTY0LFBITjJaeUJtYVd4c1BTSjNhR2wwWlhOdGIydGxJaUJ5YjJ4bFBTSnBiV2NpSUhacFpYZENiM2c5SWpBZ01DQXlOQ0F5TkNJZ2VHMXNibk05SW1oMGRIQTZMeTkzZDNjdWR6TXViM0puTHpJd01EQXZjM1puSWo0OGRHbDBiR1UrUjJsMFNIVmlQQzkwYVhSc1pUNDhjR0YwYUNCa1BTSk5NVElnTGpJNU4yTXROaTQyTXlBd0xURXlJRFV1TXpjekxURXlJREV5SURBZ05TNHpNRE1nTXk0ME16Z2dPUzQ0SURndU1qQTFJREV4TGpNNE5TNDJMakV4TXk0NE1pMHVNalU0TGpneUxTNDFOemNnTUMwdU1qZzFMUzR3TVMweExqQTBMUzR3TVRVdE1pNHdOQzB6TGpNek9DNDNNalF0TkM0d05ESXRNUzQyTVMwMExqQTBNaTB4TGpZeFF6UXVOREl5SURFNExqQTNJRE11TmpNeklERTNMamNnTXk0Mk16TWdNVGN1TjJNdE1TNHdPRGN0TGpjME5DNHdPRFF0TGpjeU9TNHdPRFF0TGpjeU9TQXhMakl3TlM0d09EUWdNUzQ0TXpnZ01TNHlNellnTVM0NE16Z2dNUzR5TXpZZ01TNHdOeUF4TGpnek5TQXlMamd3T1NBeExqTXdOU0F6TGpRNU5TNDVPVGd1TVRBNExTNDNOell1TkRFM0xURXVNekExTGpjMkxURXVOakExTFRJdU5qWTFMUzR6TFRVdU5EWTJMVEV1TXpNeUxUVXVORFkyTFRVdU9UTWdNQzB4TGpNeExqUTJOUzB5TGpNNElERXVNak0xTFRNdU1qSXRMakV6TlMwdU16QXpMUzQxTkMweExqVXlNeTR4TURVdE15NHhOellnTUNBd0lERXVNREExTFM0ek1qSWdNeTR6SURFdU1qTXVPVFl0TGpJMk55QXhMams0TFM0ek9Ua2dNeTB1TkRBMUlERXVNREl1TURBMklESXVNRFF1TVRNNElETWdMalF3TlNBeUxqSTRMVEV1TlRVeUlETXVNamcxTFRFdU1qTWdNeTR5T0RVdE1TNHlNeTQyTkRVZ01TNDJOVE11TWpRZ01pNDROek11TVRJZ015NHhOell1TnpZMUxqZzBJREV1TWpNZ01TNDVNU0F4TGpJeklETXVNaklnTUNBMExqWXhMVEl1T0RBMUlEVXVOakkxTFRVdU5EYzFJRFV1T1RJdU5ESXVNell1T0RFZ01TNHdPVFl1T0RFZ01pNHlNaUF3SURFdU5qQTJMUzR3TVRVZ01pNDRPVFl0TGpBeE5TQXpMakk0TmlBd0lDNHpNVFV1TWpFdU5qa3VPREkxTGpVM1F6SXdMalUyTlNBeU1pNHdPVElnTWpRZ01UY3VOVGt5SURJMElERXlMakk1TjJNd0xUWXVOakkzTFRVdU16Y3pMVEV5TFRFeUxURXlJaTgrUEM5emRtYysiLz48dGV4dCB4PSI0MTUiIHk9IjE0MCIgdHJhbnNmb3JtPSJzY2FsZSguMSkiIGZpbGw9IiNmYWZhZmEiIHRleHRMZW5ndGg9IjM3MCI+R2l0SHViPC90ZXh0Pjx0ZXh0IHg9Ijc2NSIgeT0iMTQwIiB0cmFuc2Zvcm09InNjYWxlKC4xKSIgZmlsbD0iI2ZhZmFmYSIgdGV4dExlbmd0aD0iMTcwIj5ia2w8L3RleHQ+PC9nPjwvc3ZnPgo=)
