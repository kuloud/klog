package klog

import (
    "testing"
)

func TestV(t *testing.T) {
    V("kuloud", "TestV", "SUCCESS")
}

func BenchmarkV(b *testing.B) {
    for i := 1; i < b.N; i++ {
        D("Benchmark", "TestV", "SUCCESS", i)
    }

}

func TestVf(t *testing.T) {
    Vf("tag", "%s", "Hello")
}

func TestD(t *testing.T) {
    D("tag", "Hello")
}

func TestDf(t *testing.T) {
    Df("tag", "%s", "Hello")
}

func TestI(t *testing.T) {
    I("tag", "Hello")
}

func TestIf(t *testing.T) {
    If("tag", "Hello = %s%s", "Kuloud", "dddd")
}

func TestW(t *testing.T) {
    W("tag", "Hello")
}

func TestWf(t *testing.T) {
    Wf("tag", "%s%s", "Kuloud", "kkk")
}

func TestE(t *testing.T) {
    E("tag", "Hello")
}

func TestEf(t *testing.T) {
    Ef("tag", "%s", "Hello")
}
