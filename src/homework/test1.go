package main

type Student struct {
	Name   string
	yuwen  float32
	shuxue float32
	waiyu  float32
	avg    float32
}

func (s *Student) stu_avg() {
	s.avg = (s.waiyu + s.yuwen + s.shuxue) / 3
}

type Class struct {
	Students []Student
	yuwen    float32
	shuxue   float32
	waiyu    float32
}

func (c *Class) class_yuwen_avg() {
	if len(c.Students) == 0 {
		return

	}
	var sum float32
	for _, stu := range c.Students {
		sum += stu.yuwen
	}
	c.yuwen = sum / float32(len(c.Students))
}
