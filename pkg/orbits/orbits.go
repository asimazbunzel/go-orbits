package orbits

import (
   "math"
	"go-orbits/pkg/io"
	"io/ioutil"

	"gonum.org/v1/gonum/stat/distuv"
	"gopkg.in/yaml.v3"
)

// structure with binary configuration
type Binary struct {
   M1 float64 `yaml:"m1"`
   M2 float64 `yaml:"m2"`
   Separation float64 `yaml:"separation"`
   Period float64 `yaml:"period"`
   
   MCO float64 `yaml:"compact_object_mass"`

   KickStrengthDistribution string `yaml:"kick_distribution"`
   KickDirection string `yaml:"kick_direction"`
   ReduceByFallback bool `yaml:"reduce_by_fallback"`

   SigmaStrength float64 `yaml:"kick_sigma"`
   NumberOfCases int `yaml:"number_of_cases"`

   W []float64
}


// read information on binary and kicks from YAML file
func (b *Binary) parseYAML (filename string) error {
  
   // read YAML data file into bytes 
   data, err := ioutil.ReadFile(filename)
   if err != nil {
      io.LogError("ORBITS - orbits.go - ParseYAML", "problem reading YAML file")
   }
   
   return yaml.Unmarshal(data, b)
}


// initialize structure Binary with the info from a binary system that will then be analyze in
// different conditions due to asymmetric momentum kicks
// it returns the Binary object
func InitBinary (filename string) Binary {

   // load binary into memory
   var binary Binary
   err := binary.parseYAML(filename)
   if err != nil {
      io.LogError("ORBITS - orbits.go - InitBinary", "unable to parse YAML file at start")
   }

   // use CGS for this
   binary.convertoCGS()

   return binary
}


// input should be in Msun / Rsun / Lsun and so on.. here we change it to CGS
func (b *Binary) convertoCGS () {

   b.M1 = b.M1 * Msun
   b.M2 = b.M2 * Msun
   b.Separation = b.Separation * Rsun
   b.Period = b.Period * 24 * 3600.0
   b.MCO = b.MCO * Msun

}


// create slices of asymmetric kicks following a given probability density function
func (b *Binary) ComputeKicks () {

   // Strength of kick based on config option
   if b.KickStrengthDistribution == "Maxwell" {

      // Maxwell distribution is just a chi-squared distribution with 3 d.o.f., k=3
      // therefore, just use inverse sampling for the chi-squared and then correct values with
      // normalization constant
      maxwell := distuv.ChiSquared{3, nil}
      for k := 1; k <= b.NumberOfCases; k++ {
         b.W = append(b.W, b.SigmaStrength * math.Sqrt(maxwell.Rand()))
      }
   } else {
      io.LogError("ORBITS - orbits.go - ComputeKicks", "unknown KickStrengthDistribution")
   }
}
