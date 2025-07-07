package githubcli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		workingDir string
		detect     packit.DetectFunc
	)

	it.Before(func() {
		var err error
		workingDir, err = os.MkdirTemp("", "working-dir")
		Expect(err).NotTo(HaveOccurred())

		parser := NewConfigParser()
		detect = Detect(parser)
	})

	it.After(func() {
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	context("when .github directory exists", func() {
		it("detects successfully", func() {
			Expect(os.MkdirAll(filepath.Join(workingDir, ".github"), 0755)).To(Succeed())

			result, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Plan.Requires).To(HaveLen(1))
			Expect(result.Plan.Requires[0].Name).To(Equal("github-cli"))
		})
	})

	context("when gh.yml exists", func() {
		it("detects successfully", func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "gh.yml"), []byte("config"), 0644)).To(Succeed())

			result, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Plan.Requires).To(HaveLen(1))
			Expect(result.Plan.Requires[0].Name).To(Equal("github-cli"))
		})
	})

	context("when package.json exists", func() {
		it("detects successfully", func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "package.json"), []byte(`{"name": "test"}`), 0644)).To(Succeed())

			result, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Plan.Requires).To(HaveLen(1))
			Expect(result.Plan.Requires[0].Name).To(Equal("github-cli"))
		})
	})

	context("when gh-requirements.txt contains github-cli", func() {
		it("detects successfully", func() {
			Expect(os.WriteFile(filepath.Join(workingDir, "gh-requirements.txt"), []byte("github-cli\nother-tool"), 0644)).To(Succeed())

			result, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Plan.Requires).To(HaveLen(1))
			Expect(result.Plan.Requires[0].Name).To(Equal("github-cli"))
		})
	})

	context("when no GitHub CLI indicators are present", func() {
		it("detects successfully but doesn't require github-cli", func() {
			result, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Plan.Requires).To(HaveLen(1))
			Expect(result.Plan.Requires[0].Name).To(Equal("github-cli"))
		})
	})
} 