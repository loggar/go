# time

## toTime

```
func toTime(ms int64) time.Time {
	sec, nsec := ms/1000, (ms%1000)*1000
	return time.Unix(sec, nsec)
}
```

```
timeValue, err := time.Parse("2006-01-02 15:04:05", "2023-02-07 20:30:00")
if err != nil {
	// handle the error
} else {
	c.certVerifyOption.CurrentTime = timeValue
}
```
