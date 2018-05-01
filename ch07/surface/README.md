# Surface Web Application

## Run

```
go run ch07/surface/surface.go
```

## Test

```
http://localhost:8000/plot?expr=sin(-x)*pow(1.5,-r)
```

```
http://localhost:8000/plot?expr=pow(2,sin(y))*pow(2,sin(x))/12
```

```
http://localhost:8000/plot?expr=sin(x*y/10)/10
```
