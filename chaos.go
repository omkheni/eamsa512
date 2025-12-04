// chaos.go
package main

import (
    "math"
    "math/rand"
    "time"
)

// Vector3 represents a 3D vector for Lorenz system
type Vector3 struct {
    X, Y, Z float64
}

// Vector5 represents a 5D vector for hyperchaotic system
type Vector5 struct {
    M, N, P, R, Q float64
}

// Lorenz parameters
const (
    sigma = 10.0
    rho   = 28.0
    beta  = 8.0 / 3.0
)

// Hyperchaotic parameters
const (
    a = 30.0
    b = 11.0
    c = 90.0
)

// RK4 step for Lorenz
func lorenzRK4(v Vector3, dt float64) Vector3 {
    // Implement vectorized RK4 for Lorenz
    k1 := lorenzDeriv(v)
    k2 := lorenzDeriv(Vector3{
        X: v.X + 0.5*dt*k1.X,
        Y: v.Y + 0.5*dt*k1.Y,
        Z: v.Z + 0.5*dt*k1.Z,
    })
    k3 := lorenzDeriv(Vector3{
        X: v.X + 0.5*dt*k2.X,
        Y: v.Y + 0.5*dt*k2.Y,
        Z: v.Z + 0.5*dt*k2.Z,
    })
    k4 := lorenzDeriv(Vector3{
        X: v.X + dt*k3.X,
        Y: v.Y + dt*k3.Y,
        Z: v.Z + dt*k3.Z,
    })

    return Vector3{
        X: v.X + (dt/6.0)*(k1.X+2.0*k2.X+2.0*k3.X+k4.X),
        Y: v.Y + (dt/6.0)*(k1.Y+2.0*k2.Y+2.0*k3.Y+k4.Y),
        Z: v.Z + (dt/6.0)*(k1.Z+2.0*k2.Z+2.0*k3.Z+k4.Z),
    }
}

func lorenzDeriv(v Vector3) Vector3 {
    return Vector3{
        X: sigma * (v.Y - v.X),
        Y: v.X*(rho - v.Z) - v.Y,
        Z: v.X*v.Y - beta*v.Z,
    }
}

// RK4 step for hyperchaotic system
func hyperchaoticRK4(v Vector5, dt float64) Vector5 {
    k1 := hyperchaoticDeriv(v)
    k2 := hyperchaoticDeriv(Vector5{
        M: v.M + 0.5*dt*k1.M,
        N: v.N + 0.5*dt*k1.N,
        P: v.P + 0.5*dt*k1.P,
        R: v.R + 0.5*dt*k1.R,
        Q: v.Q + 0.5*dt*k1.Q,
    })
    k3 := hyperchaoticDeriv(Vector5{
        M: v.M + 0.5*dt*k2.M,
        N: v.N + 0.5*dt*k2.N,
        P: v.P + 0.5*dt*k2.P,
        R: v.R + 0.5*dt*k2.R,
        Q: v.Q + 0.5*dt*k2.Q,
    })
    k4 := hyperchaoticDeriv(Vector5{
        M: v.M + dt*k3.M,
        N: v.N + dt*k3.N,
        P: v.P + dt*k3.P,
        R: v.R + dt*k3.R,
        Q: v.Q + dt*k3.Q,
    })

    return Vector5{
        M: v.M + (dt/6.0)*(k1.M+2.0*k2.M+2.0*k3.M+k4.M),
        N: v.N + (dt/6.0)*(k1.N+2.0*k2.N+2.0*k3.N+k4.N),
        P: v.P + (dt/6.0)*(k1.P+2.0*k2.P+2.0*k3.P+k4.P),
        R: v.R + (dt/6.0)*(k1.R+2.0*k2.R+2.0*k3.R+k4.R),
        Q: v.Q + (dt/6.0)*(k1.Q+2.0*k2.Q+2.0*k3.Q+k4.Q),
    }
}

func lorenzDeriv(v Vector3) Vector3 {
    return Vector3{
        X: sigma * (v.Y - v.X),
        Y: v.X*(rho - v.Z) - v.Y,
        Z: v.X*v.Y - beta*v.Z,
    }
}

func hyperchaoticDeriv(v Vector5) Vector5 {
    return Vector5{
        M: a*(v.N - v.M),
        N: v.M*(b - v.P) - v.N + v.Q,
        P: v.M*v.N - c*v.P,
        R: v.N*v.P - v.R,
        Q: v.R - v.Q,
    }
}

// Initialize chaos systems with seed
func initChaos(seed int64) (Vector3, Vector5) {
    rand.Seed(seed)
    vLorenz := Vector3{
        X: rand.Float64() * 20 - 10,
        Y: rand.Float64() * 20 - 10,
        Z: rand.Float64() * 20 - 10,
    }
    vHyper := Vector5{
        M: rand.Float64() * 30,
        N: rand.Float64() * 30,
        P: rand.Float64() * 30,
        R: rand.Float64() * 30,
        Q: rand.Float64() * 30,
    }
    return vLorenz, vHyper
}

// Generate chaos keys
func generateChaosKeys(seed int64, steps int, dt float64) [11][]byte {
    vLorenz, vHyper := initChaos(seed)
    var keys [11][]byte
    for i := 0; i < steps; i++ {
        vLorenz = lorenzRK4(vLorenz, dt)
        vHyper = hyperchaoticRK4(vHyper, dt)
        // Map states to bytes
        keys[0] = append(keys[0], float64ToBytes(vLorenz.X)...)
        keys[1] = append(keys[1], float64ToBytes(vLorenz.Y)...)
        keys[2] = append(keys[2], float64ToBytes(vLorenz.Z)...)
        keys[3] = append(keys[3], float64ToBytes(vHyper.M)...)
        keys[4] = append(keys[4], float64ToBytes(vHyper.N)...)
        keys[5] = append(keys[5], float64ToBytes(vHyper.P)...)
        keys[6] = append(keys[6], float64ToBytes(vHyper.R)...)
        keys[7] = append(keys[7], float64ToBytes(vHyper.Q)...)
        // Additional states can be added as needed
    }
    // Hash or normalize as needed
    return keys
}

func float64ToBytes(f float64) []byte {
    return []byte{
        byte(math.Float64bits(f) >> 56),
        byte(math.Float64bits(f) >> 48),
        byte(math.Float64bits(f) >> 40),
        byte(math.Float64bits(f) >> 32),
        byte(math.Float64bits(f) >> 24),
        byte(math.Float64bits(f) >> 16),
        byte(math.Float64bits(f) >> 8),
        byte(math.Float64bits(f)),
    }
}
